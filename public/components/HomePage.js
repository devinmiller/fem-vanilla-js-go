import { API } from "../services/api.js";

export class HomePage extends HTMLElement {

  async render() {
    const topResult = await API.getTopMovies();
    renderMoviesInList(topResult.movies, document.querySelector("#top-10 ul"));

    const randomResult = await API.getRandomMovies();
    renderMoviesInList(randomResult.movies, document.querySelector("#random ul"));

    function renderMoviesInList(movies, ul) {
      ul.innerHTML = "";
      movies.forEach(movie => {
        const li = document.createElement("li");
        li.textContent = movie.title;
        ul.appendChild(li);
      });
    }
  }

  connectedCallback() {
    const template = document.getElementById("template-home");
    const content = template.content.cloneNode(true)
    this.appendChild(content)

    this.render()
  }
}

customElements.define("home-page", HomePage);
