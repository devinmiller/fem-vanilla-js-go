import { API } from "./services/api.js";

import "./components/HomePage.js"
import "./components/MovieDetailsPage.js"
import "./components/AnimatedLoading.js"
import "./components/YoutubeEmbed.js"
import Router from "./services/Router.js";

window.app = {
  Router,
  search: (event) => {
    event.preventDefault();
    const q = document.querySelector("input[type=search]").value;
  },
  api: API
};

window.addEventListener("DOMContentLoaded", () => {
  app.Router.init();
});


