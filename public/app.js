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
  api: API,
  showError: (message = "There was an error loading the page", goToHome = true) => {
    document.querySelector("#alert-modal").showModal();
    document.querySelector("#alert-modal p").textContents = message;
    if (goToHome) {
      app.Router.go("/");
    }
  },
  closeError: () => {
    document.getElementById("alert-modal").close();
  }
};

window.addEventListener("DOMContentLoaded", () => {
  app.Router.init();
});


