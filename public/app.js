import { API } from "./services/api.js";

import "./components/HomePage.js"
import "./components/MovieDetailsPage.js"
import "./components/AnimatedLoading.js"
import "./components/YoutubeEmbed.js"

window.app = {
  search: (event) => {
    event.preventDefault();
    const q = document.querySelector("input[type=search]").value;
  },
  api: API
};

