<script lang="ts">
	import { onMount } from "svelte";

    interface ImageData {
        id: number;
        url: string;
        elo: number;
        k_factor: number;
        rounds: number;
    }

    let Image1: ImageData = $state({
        id: 6,
        url: "",
        elo: 1400,
        k_factor: 40,
        rounds: 0,
    });

    let Image2: ImageData = $state({
        id: 9,
        url: "",
        elo: 1400,
        k_factor: 40,
        rounds: 0,
    });

    // State for selected image
    let selected = $state(0);
    // State for image URLs from API
    let loading = $state(true);
    const IMAGE_HEIGHT = '250px';

    const HOST = import.meta.env.DEV ? "http://localhost:8080" : "https://footyrate.onrender.com";

    // API URL
    const API_URL = HOST + '/api/random-images';
    const API_RESULT = HOST + '/api/result';

    async function submit_vote(winnerID: number, loserID: number) {
        try {
            const response = await fetch(API_RESULT, {
                method: 'POST',
                headers: {
                    'Content-Type' : 'application/json',
                },
                body: JSON.stringify({
                    winner_ID: winnerID,
                    loser_ID: loserID,
                })
            })

            if (!response.ok) {
                throw new Error('Error submitting vote');
            }
        } catch(err) {
            console.error("Error occurred submitting vote", err);
        }
    }
    
    // Function to fetch random images
    async function fetchRandomImages() {
        loading = true;
        try {
            const response = await fetch(API_URL);
            if (!response.ok) {
                throw new Error('Failed to fetch images');
            }
            const data = await response.json();
            Image1 = data.image1;
            Image2 = data.image2;

            // Reset selection when new images are loaded
            selected = 0;
        } catch (error) {
            console.error('Error fetching images:', error);
        } finally {
            loading = false;
        }
    }

    function handle_vote(n: number) {
        if (n == 1) {
            submit_vote(Image1.id, Image2.id)
        } else {
            submit_vote(Image2.id, Image1.id)
        }
    }
    function handle_selection (n: number) {
        selected = n;

        handle_vote(n);
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
        flex-wrap: wrap;
        gap: 30px;
        width: 100%;

        /* display: grid;
        justify-content: space-around;
        align-items: center;
        width: 95%;
        background-color: #C9E4E7;
        border-radius: 10px;
        padding: 20px; */
    }
    container button {
        flex: 1;
        min-width: 300px; /* Adjust this value based on your needs */
        border: none;
        padding: 0;
        background: none;
        transition: 0.3s;
        cursor: pointer;
    }
    
    container img {
        width: 100%;
        height: auto;
        max-height: var(--IMAGE_HEIGHT);
        object-fit: contain;
    }

    @media (max-width: 768px) {
        container button {
        flex-basis: 100%;
        }
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
                <img src={Image1.url} height={IMAGE_HEIGHT} alt="First">
                <!-- <img src={Image1.url} alt="First"> -->
            </button>
            <button onclick={() => handle_selection(2)}>
                <img src={Image2.url} height={IMAGE_HEIGHT} alt="Second">
                <!-- <img src={Image2.url} alt="Second"> -->
            </button>
        </container>
        
        {#if selected != 0}
            <p class="result">Image {selected} selected</p>
        {:else}
            <p class="result">Select an image</p>
        {/if}
    {/if}
</main>