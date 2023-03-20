<script lang="ts">
  import { ListDir, OpenFile } from "../wailsjs/go/main/App.js";
  let pathTokens = [];
  let hidden = false;
  let rowSelected = 0;
  let filterString = "";
  $: path = pathTokens.length > 0 ? `/${pathTokens.join("/")}` : "/";
  $: dirListing = ListDir(path, hidden);

  async function filteredList() {
    return (await dirListing).filter((d) =>
      d.Name.toLowerCase().includes(filterString.toLowerCase())
    );
  }

  async function handleSelect() {
    const selected = (await filteredList())[rowSelected];
    if (selected.IsDir) {
      pathTokens = [...pathTokens, selected.Name];
      reset();
    } else {
      OpenFile(`${path}/${selected.Name}`);
    }
  }

  function reset() {
    rowSelected = 0;
    filterString = "";
  }

  function goUpOneLevel() {
    pathTokens = pathTokens.slice(0, -1);
    reset();
  }

  async function handleKeyDown(e: KeyboardEvent) {
    switch (e.key) {
      case "Enter":
        handleSelect();
        break;
      case "ArrowDown":
        if (rowSelected < (await filteredList()).length - 1) {
          rowSelected++;
        }
        break;
      case "ArrowUp":
        if (e.metaKey) {
          goUpOneLevel();
        } else {
          rowSelected > 0 && rowSelected--;
        }
        break;
      case "Backspace":
        if (filterString.length > 0) {
          filterString = filterString.slice(0, -1);
        } else {
          goUpOneLevel();
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
          pathTokens = ["Users", "brent"];
          reset();
        }
        break;
      case "Escape":
        filterString = "";
        break;
      default:
        if (e.key.length === 1 && !e.metaKey && !e.altKey && !e.ctrlKey) {
          if (filterString.length === 0 && e.key === " ") {
            break;
          }
          filterString += e.key;
        }
    }
  }
</script>

<svelte:window on:keydown|preventDefault={handleKeyDown} />

<main>
  <div>
    <div class="path">
      {#if pathTokens.length === 0}
        {` / `}
      {/if}
      {#each pathTokens as token, idx}
        {` / `}
        <!-- svelte-ignore a11y-click-events-have-key-events -->
        <span
          class="pathToken"
          on:click|preventDefault={() => {
            pathTokens = pathTokens.slice(0, idx + 1);
            reset();
          }}>{token}</span
        >
      {/each}
    </div>
    <table>
      <tr><td class="li-header"><strong>File</strong></td></tr>
      {#await dirListing then dirs}
        {#each dirs.filter( (d) => d.Name.toLowerCase().includes(filterString.toLowerCase()) ) as dir, i}
          <tr
            class={i === rowSelected ? "selected" : `row-${i % 2}`}
            on:click={() => (rowSelected = i)}
          >
            <!-- svelte-ignore a11y-click-events-have-key-events -->
            <td width="400px">
              <span class={dir.IsDir ? "dir" : "file"}>{dir.Name}</span>
            </td>
            <td width="100px">
              {#if !dir.IsDir}
                {dir.Size}
              {/if}
            </td>
            <td>
              {dir.ModTime}
            </td>
          </tr>
        {/each}
      {/await}
    </table>
  </div>
</main>

<style>
  .path {
    padding: 10px;
  }
  .pathToken {
    margin: 2px;
    padding: 6px;
    background-color: rgba(83, 93, 110, 1);
    border-radius: 0.25em;
    cursor: default;
  }
  span.dir {
    font-weight: 400;
    color: cornsilk;
  }
  td {
    list-style: none;
    user-select: none;
    -webkit-user-select: none;
    cursor: default;
    padding: 0.2em 2em;
  }
  tr.selected {
    background-color: rgba(83, 93, 110, 1);
  }
  tr.row-0 {
    background-color: rgba(35, 45, 62, 1);
  }
</style>
