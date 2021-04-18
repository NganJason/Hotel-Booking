let notify = function (msg, msgType) {
  notie.alert({
    type: msgType,
    text: msg,
  });
};

let alertMsg = function (title, msg, icon) {
  swal({
    title: title,
    text: msg,
    icon: icon,
  });
};

let toast = function (c) {
  const { msg = "", icon = "success", position = "top-end" } = c;

  const Toast = Swal.mixin({
    toast: true,
    title: msg,
    position: position,
    icon: icon,
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: (toast) => {
      toast.addEventListener("mouseenter", Swal.stopTimer);
      toast.addEventListener("mouseleave", Swal.resumeTimer);
    },
  });

  Toast.fire({});
};

let success = function (c) {
  const { msg = "", title = "", footer = "" } = c;

  Swal.fire({
    icon: "success",
    title: title,
    text: msg,
    footer: footer,
  });
};

let error = function (c) {
  const { msg = "", title = "", footer = "" } = c;

  Swal.fire({
    icon: "error",
    title: title,
    text: msg,
    footer: footer,
  });
};

let custom = async function (c) {
  const { msg = "", title = "" } = c;

  const { value: result } = await Swal.fire({
    title: title,
    html: msg,
    backdrop: false,
    focusConfirm: false,
    showCancelButton: true,
    willOpen: () => {
      if (c.willOpen != undefined) {
        c.willOpen();
      }
    },
    didOpen: () => {
      if (c.didOpen != undefined) {
        c.didOpen();
      }
    },
    preConfirm: () => {
      return [
        document.getElementById("start").value,
        document.getElementById("end").value,
      ];
    },
  });

  if (result) {
    if (result.dismiss !== Swal.DismissReason.cancel) {
      if (result.value !== "") {
        if (c.callback !== undefined) {
          c.callback(result);
        } else {
          c.callback(false);
        }
      } else {
        c.callback(false);
      }
    }
  }
};

let Prompt = () => {
  return {
    notify: notify,
    alertMsg: alertMsg,
    toast: toast,
    success: success,
    error: error,
    custom: custom,
  };
};

export let attention = Prompt();
