export function handleKeyPress(e: KeyboardEvent) {
  switch (e.key) {
    case "Enter":
      handleSelect();
      break;
    case "ArrowDown":
      rowSelected++;
      break;
    case "ArrowUp":
      if (e.metaKey) {
        if (path.lastIndexOf("/") !== 0) {
          path = path.slice(0, path.lastIndexOf("/"));
          reset();
        } else {
          path = "/";
          reset();
        }
      } else {
        rowSelected > 0 && rowSelected--;
      }
      break;
    case "Home":
      rowSelected = 0;
      break;
    case "End":
      rowSelected = (await dirListing).length - 1;
      break;
    case ".":
      if (e.metaKey && e.shiftKey) {
        hidden = !hidden;
      }
      break;
    case "h":
      if (e.metaKey && e.shiftKey) {
        path = "/Users/brent";
        reset();
        break;
      }
    case "Escape":
      filterString = "";
      break;
    default:
      if (e.key.length === 1) filterString += e.key;
  }
}
