export class MovieItemComponent extends HTMLElement {
  constructor(movie) {
    super();
    this.movie = movie;
  }

  connectedCallback() {
    this.innerHTML = `
    <a href="#" class="movie-details-link">
      <article>
        <img src="${this.movie.poster_url}" alt="${this.movie.title} Poster">
        <p>${this.movie.title} (${this.movie.release_year})</p>
      </article>
    </a>
    `;

    const link = this.querySelector(".movie-details-link");
    // prevent default and navigate on click
    link.addEventListener("click", (event) => {
      event.preventDefault();
      const url = `/movies/${this.movie.id}`;
      app.Router.go(url);
    });

  }
}
customElements.define("movie-item", MovieItemComponent);
