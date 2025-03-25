<script lang="ts">
	import { onMount } from "svelte";

    // State for selected image
    let selected = $state(0);
    // State for image URLs from API
    let img1 = $state('');
    let img2 = $state('');
    let loading = $state(true);
    const IMAGE_HEIGHT = '300px';
    
    // API URL
    const API_URL = 'http://localhost:8080/api/random-images';
    
    // Function to fetch random images
    async function fetchRandomImages() {
        loading = true;
        try {
            const response = await fetch(API_URL);
            if (!response.ok) {
                throw new Error('Failed to fetch images');
            }
            const data = await response.json();
            img1 = data.image1;
            img2 = data.image2;
            // Reset selection when new images are loaded
            selected = 0;
        } catch (error) {
            console.error('Error fetching images:', error);
        } finally {
            loading = false;
        }
    }

    function handle_selection (n: number) {
        selected = n;

        fetchRandomImages();
    }
    
    // Load initial images
    onMount(() => {
        fetchRandomImages();
    });
</script>

<style>
    * {
        padding: 0;
        margin: 0;
    }
    main {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
    }
    container {
        display: flex;
        justify-content: space-around;
        align-items: center;
        width: 95%;
        background-color: #C9E4E7;
        border-radius: 10px;
        padding: 20px;
    }
    button {
        width: 50%;
        margin: 20px;
        border: none;
        transition: 0.3s;
        background: none;
        cursor: pointer;
    }
    button:hover {
        transform: scale(1.05);
    }
    button:active {
        transform: translateY(4px);
    }
    .result {
        padding-top: 20px;
        margin-bottom: 20px;
    }
    .loading {
        height: 300px;
        display: flex;
        align-items: center;
        justify-content: center;
    }
</style>

<main>
    <h1>Choose</h1>
    
    {#if loading}
        <div class="loading">Loading images...</div>
    {:else}
        <container>
            <button onclick={() => handle_selection(1)}>
                <img src={img1} height={IMAGE_HEIGHT} alt="First">
            </button>
            <button onclick={() => handle_selection(2)}>
                <img src={img2} height={IMAGE_HEIGHT} alt="Second">
            </button>
        </container>
        
        {#if selected != 0}
            <p class="result">Image {selected} selected</p>
        {:else}
            <p class="result">Select an image</p>
        {/if}
    {/if}
</main>