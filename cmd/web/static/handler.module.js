import { attention } from "./alert.module.js";
import { checkAvailabilityStaticForm } from "./static.module.js";

export let checkAvailabilityHandler = function () {
  let html = checkAvailabilityStaticForm;
  attention.custom({
    title: "Choose your dates",
    msg: html,
    callback: function (result) {
      let form = document.getElementById("check-availability-form");
      let formData = new FormData(form);

      fetch("/search-availability-json", {
        method: "post",
        body: formData,
      })
        .then((response) => response.json())
        .then((data) => {
          console.log(data);
        });
    },
    willOpen: () => {
      const elem = document.getElementById("reservation-dates-modal");
      const rp = new DateRangePicker(elem, {
        format: "yyyy-mm-dd",
        showOnFocus: true,
      });
    },
    didOpen: () => {
      document.getElementById("start").removeAttribute("disabled");
      document.getElementById("end").removeAttribute("disabled");
    },
  });
};
