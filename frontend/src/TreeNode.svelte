<script context="module">
  export let location;

  let isOpen = false;

  // Toggle child elements
  function toggle() {
    isOpen = !isOpen;
  }
</script>

<li>
  <button type="button" class="toggle" on:click={toggle}>
    {location.name}
    {isOpen ? "▼" : "▶"}
  </button>

  {#if isOpen}
    <ul class="tree">
      {#each getChildren(location.id) as childLocation}
        <svelte:self {childLocation} />
      {/each}
      {#each getItems(location.id) as item}
        <button type="button" on:click={() => selectItem(item)}
          >{item.name}</button
        >
      {/each}
    </ul>
  {/if}
</li>

<!-- Item Details Section -->
{#if selectedItem}
  <div class="content">
    <h2>{selectedItem.name}</h2>
    <p>Quantity: {selectedItem.quantity}</p>
    <p>Category: {selectedItem.category}</p>
    <p>Price: ₹{selectedItem.price}</p>
    <p>Status: {selectedItem.status}</p>
    <img src={selectedItem.image_url} alt={selectedItem.name} width="100" />
  </div>
{/if}
