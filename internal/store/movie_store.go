package store

import (
	"database/sql"
	"fmt"

	"github.com/devinmiller/fem-vanilla-js-go/internal/models"
)

const defaultLimit = 20

type PostgresMovieStore struct {
	db *sql.DB
}

func NewPostgresMovieStore(db *sql.DB) *PostgresMovieStore {
	return &PostgresMovieStore{db: db}
}

type MovieStore interface {
	GetTopMovies() ([]models.Movie, error)
	GetRandomMovies() ([]models.Movie, error)
	GetMovieById(id int) (models.Movie, error)
	// SearchMoviesByName(name string) ([]models.Movie, error)
	// GetAllGenres() ([]models.Genre, error)
}

func (pg *PostgresMovieStore) getMovies(query string) ([]models.Movie, error) {
	rows, err := pg.db.Query(query, defaultLimit)
	if err != nil {
		return nil, fmt.Errorf("Failed to query db for movies: %w", err)
	}
	defer rows.Close()

	var movies []models.Movie
	for rows.Next() {
		var m models.Movie
		err = rows.Scan(
			&m.ID,
			&m.TMDB_ID,
			&m.Title,
			&m.Tagline,
			&m.ReleaseYear,
			&m.Overview,
			&m.Score,
			&m.Popularity,
			&m.Language,
			&m.PosterURL,
			&m.TrailerURL,
		)
		if err != nil {
			return nil, fmt.Errorf("Failed to scan rows for movies: %w", err)
		}

		err = pg.getMovieRelations(&m)
		if err != nil {
			return nil, fmt.Errorf("Failed to get relations for movies: %w", err)
		}

		movies = append(movies, m)
	}

	return movies, nil
}

func (pg *PostgresMovieStore) getMovieRelations(m *models.Movie) error {
	// get genres
	genreQuery := `
	SELECT g.id, g.name
	FROM genres g
	JOIN movie_genres mg on g.id = mg.genre_id
	WHERE mg.movie_id = $1
	`

	genreRows, err := pg.db.Query(genreQuery, m.ID)
	if err != nil {
		return fmt.Errorf("Failed to query db for genres: %w", err)
	}
	defer genreRows.Close()

	for genreRows.Next() {
		var genre models.Genre
		err = genreRows.Scan(&genre.ID, &genre.Name)
		if err != nil {
			return fmt.Errorf("Failed to scan rows for genres: %w", err)
		}
		m.Genres = append(m.Genres, genre)
	}

	// get actors
	actorQuery := `
	SELECT a.ID, a.first_name, a.last_name, a.image_url
	FROM actors a
	JOIN movie_cast mc on a.id = mc.actor_id
	WHERE mc.movie_id = $1
	`

	actorRows, err := pg.db.Query(actorQuery, m.ID)
	if err != nil {
		return fmt.Errorf("Failed to query db for actors: %w", err)
	}
	defer actorRows.Close()

	for actorRows.Next() {
		var actor models.Actor
		err = actorRows.Scan(&actor.ID, &actor.FirstName, &actor.LastName, &actor.ImageURL)
		if err != nil {
			return fmt.Errorf("Failed to scan rows for actors: %w", err)
		}
		m.Casting = append(m.Casting, actor)
	}

	keywordsQuery := `
	SELECT k.word
	FROM keywords k
	JOIN movie_keywords mk ON k.id = mk.keyword_id
	WHERE mk.movie_id = $1
	`

	keywordRows, err := pg.db.Query(keywordsQuery, m.ID)
	if err != nil {
		return fmt.Errorf("Failed to query db for keywords: %w", err)
	}
	defer keywordRows.Close()

	for keywordRows.Next() {
		var keyword string
		err = keywordRows.Scan(&keyword)
		if err != nil {
			return fmt.Errorf("Failed to scan rows for keywords: %w", err)
		}
		m.Keywords = append(m.Keywords, keyword)
	}

	return nil
}

func (pg *PostgresMovieStore) GetTopMovies() ([]models.Movie, error) {
	query := `
	SELECT id, tmdb_id, title, tagline, release_year, overview, score, popularity, language, poster_url, trailer_url
	FROM movies
	ORDER BY popularity DESC
	LIMIT $1
	`

	movies, err := pg.getMovies(query)
	if err != nil {
		return nil, fmt.Errorf("GetTopMovies: %w", err)
	}

	return movies, nil
}

func (pg *PostgresMovieStore) GetRandomMovies() ([]models.Movie, error) {
	query := `
	SELECT id, tmdb_id, title, tagline, release_year, overview, score, popularity, language, poster_url, trailer_url
	FROM movies
	ORDER BY random()
	LIMIT $1
	`

	movies, err := pg.getMovies(query)
	if err != nil {
		return nil, fmt.Errorf("GetRandomMovies: %w", err)
	}

	return movies, nil
}

func (pg *PostgresMovieStore) GetMovieById(id int) (models.Movie, error) {
	query := `
	SELECT id, tmdb_id, title, tagline, release_year, overview, score, popularity, language, poster_url, trailer_url
	FROM movies
	WHERE id = $1
	`

	row := pg.db.QueryRow(query, id)

	var m models.Movie
	err := row.Scan(
		&m.ID,
		&m.TMDB_ID,
		&m.Title,
		&m.Tagline,
		&m.ReleaseYear,
		&m.Overview,
		&m.Score,
		&m.Popularity,
		&m.Language,
		&m.PosterURL,
		&m.TrailerURL,
	)

	if err == sql.ErrNoRows {
		return models.Movie{}, fmt.Errorf("Failed to get movie for id: %d", id)
	}

	if err != nil {
		return models.Movie{}, fmt.Errorf("")
	}

	err = pg.getMovieRelations(&m)
	if err != nil {
	}

	return m, nil
}
