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
      let payload = {
        start_date: formData.getAll("start_date")[0],
        end_date: formData.getAll("end_date")[0],
      };

      fetch("/search-availability-json", {
        method: "POST",
        body: JSON.stringify(payload),
      })
        .then((response) => response.json())
        .then((data) => {
          if (data.ok) {
            attention.custom({
              icon: "success",
              showConfirmButton: false,
              msg:
                "<p> Room is available!</p>" +
                '<p><a href="/book-room?id=' +
                data.room_id +
                "&s=" +
                data.start_date +
                "&e=" +
                data.end_date +
                '" class="btn btn-primary">Book now!</a></p>',
            });
          } else {
            attention.error({
              msg: "No availability",
            });
          }
        });
    },
    willOpen: () => {
      const elem = document.getElementById("reservation-dates-modal");
      const rp = new DateRangePicker(elem, {
        format: "yyyy-mm-dd",
        showOnFocus: true,
        minDate: new Date(),
      });
    },
    didOpen: () => {
      document.getElementById("start").removeAttribute("disabled");
      document.getElementById("end").removeAttribute("disabled");
    },
  });
};
