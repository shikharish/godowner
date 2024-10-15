<script>
  export let children;
  export let getChildren;
  export let getItems;
  export let selectItem;

  let isOpen = false;

  function toggle() {
    isOpen = !isOpen;
  }
</script>

<div>
  <button class="toggle" on:click={toggle}>
    {isOpen ? "▼" : "▶"}
    {children.name}
  </button>

  {#if isOpen}
    <div class="child-container">
      {#each getChildren(children.id) as child}
        <svelte:self children={child} {getChildren} {getItems} {selectItem} />
      {/each}

      {#each getItems(children.id) as item}
        <div class="child-item" on:click={() => selectItem(item)}>
          {item.name}
        </div>
      {/each}
    </div>
  {/if}
</div>

<style>
  .toggle {
    cursor: pointer;
    display: flex;
    align-items: center;
    gap: 0.5rem;
    background-color: transparent;
    color: white;
    border: none;
    padding: 0.5rem 1rem;
    border-radius: 8px;
    transition:
      background-color 0.3s,
      transform 0.2s;
  }

  .toggle:hover {
    background-color: #3700b3;
    transform: scale(1.05);
  }

  .child-container {
    margin-left: 2rem;
  }

  .child-item {
    background-color: #715e9c;
    padding: 0.5rem;
    text-align: left;
    margin-top: 0.5rem;
    /* padding-left: 2rem; */
    border-radius: 6px;
    cursor: pointer;
    transition: background-color 0.3s;
  }

  .child-item:hover {
    background-color: #3700b3;
    color: white;
  }
</style>
