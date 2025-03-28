<script lang="ts">
	import { onMount } from "svelte";

    interface ImageData {
        id: number;
        url: string;
        elo: number;
    }

    let Image1: ImageData = $state({
        id: 6,
        url: "",
        elo: 1400,
    });

    let Image2: ImageData = $state({
        id: 9,
        url: "",
        elo: 1400,
    });

    // State for selected image
    let selected = $state(0);
    // State for image URLs from API
    let loading = $state(true);
    const IMAGE_HEIGHT = '300px';
    
    // API URL
    const API_URL = 'https://footyrate.onrender.com/api/random-images';
    const API_RESULT = 'https://footyrate.onrender.com/api/result';

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

            console.log(winnerID, loserID)

            if (!response.ok) {
                throw new Error('Error submitting vote');
            }

            const data = await response.json();
            console.log("This ID WON: ", data.winner_ID)
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

            console.log("First ID", Image1.id)
            console.log("Second ID", Image2.id)

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
                <img src={Image1.url} height={IMAGE_HEIGHT} alt="First">
            </button>
            <button onclick={() => handle_selection(2)}>
                <img src={Image2.url} height={IMAGE_HEIGHT} alt="Second">
            </button>
        </container>
        
        {#if selected != 0}
            <p class="result">Image {selected} selected</p>
        {:else}
            <p class="result">Select an image</p>
        {/if}
    {/if}
</main>