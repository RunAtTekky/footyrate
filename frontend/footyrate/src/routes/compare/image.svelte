<script lang="ts">
    const { image, handler, side} = $props();
</script>

{#if side === "left"}
    <button class="left" onclick={() => handler(1)}>
        <img src={image.url} alt="First">
    </button>
{:else}
    <button class="right" onclick={() => handler(2)}>
        <img src={image.url} alt="Second">
    </button>
{/if}


<style>

button {
    display: flex;
    width: 50%;
    height: 100vh;
    justify-content: center;
    min-width: 300px; /* Adjust this value based on your needs */
    border: none;
    padding: 0;
    background: none;
    transition: 0.3s;
    cursor: pointer;
}

img {
    width: 100%;
    height: 100%;
    border: none;
    max-height: var(--IMAGE_HEIGHT);
    object-fit: cover;
    transition: 0.3s;
}

img:active {
    transform: scale(0.98);
}

@media (max-width: 640px) {
    button {
        flex-basis: 100%;
        height: 50vh;
    }

    button:hover {
        transform: scale(1);
    }
    img {
        height: 100%;
        width: 100%;
        object-fit: cover;
    }

}

</style>