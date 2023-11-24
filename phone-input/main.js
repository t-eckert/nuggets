function setupInput(o) {
  o.addEventListener("keyDown", (e) => {
    console.log(e);
  });

  console.log(o);
}

document.querySelectorAll("input").forEach((o) => setupInput(o));
