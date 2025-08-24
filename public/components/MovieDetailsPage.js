import { API } from "../services/api.js";

export class MovieDetailsPage extends HTMLElement {
  id = null;
  movie = null;

  async render() {
    try {
      const result = await API.getMovieById(this.id);
      this.movie = result.movie;
      console.log(this.movie);
    } catch (e) {
      // TODO: alert user
      console.log("error", e);
      return;
    }
    const template = document.getElementById("template-movie-details");
    const content = template.content.cloneNode(true);
    this.appendChild(content);

    this.querySelector("h2").textContent = this.movie.title;
    this.querySelector("h3").textContent = this.movie.tagline;
    this.querySelector("img").src = this.movie.poster_url;
    this.querySelector("#trailer").dataset.url = this.movie.trailer_url;
    this.querySelector("#overview").textContent = this.movie.overview;
    this.querySelector("#metadata").innerHTML = `
      <dt>Release Date</dt>
      <dd>${this.movie.release_year}</dd>                        
      <dt>Score</dt>
      <dd>${this.movie.score} / 10</dd>                        
      <dt>Original languae</dt>
      <dd>${this.movie.language}</dd>  
    `;

    const ulGenres = this.querySelector("#genres");
    ulGenres.innerHTML = "";
    this.movie.genres.forEach(genre => {
      const li = document.createElement("li");
      li.textContent = genre.name;
      ulGenres.appendChild(li);
    });

    const ulCast = this.querySelector("#cast");
    ulCast.innerHTML = "";
    this.movie.casting.forEach(actor => {
      const li = document.createElement("li");
      li.innerHTML = `
        <img src="${actor.image_url ?? '/images/generic_actor.jpg'}" alt="Picture of ${actor.last_name}">
        <p>${actor.first_name} ${actor.last_name}</p>
      `;
      ulCast.appendChild(li);
    });
  }

  connectedCallback() {
    this.id = 15 // TODO: this.dataset.id; // data-id
    this.render();
  }
}
customElements.define("movie-details-page", MovieDetailsPage);
