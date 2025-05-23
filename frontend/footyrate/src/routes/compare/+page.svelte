<script lang="ts">
	import { onMount } from "svelte";
	import Pop from "./pop.svelte";
	import Image from "./image.svelte";

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
    let show_result = $state(false);
    let winner_at_left = $state(true);
    let winner_elo_change = $state(0);
    let loser_elo_change = $state(0);
    const IMAGE_HEIGHT = '100px';

    const HOST = import.meta.env.DEV ? "http://localhost:8080" : "https://footyrate.onrender.com";

    // API URL
    const API_URL = HOST + '/api/random-images';
    const API_RESULT = HOST + '/api/result';

    async function submit_vote(winner: ImageData, loser: ImageData) {
        var current_winner_elo = winner.elo;
        var current_loser_elo = loser.elo;
        try {
            const response = await fetch(API_RESULT, {
                method: 'POST',
                headers: {
                    'Content-Type' : 'application/json',
                },
                body: JSON.stringify({
                    winner_ID: winner.id,
                    loser_ID: loser.id,
                })
            })

            if (!response.ok) {
                throw new Error('Error submitting vote');
            }

            const data = await response.json();

            winner_elo_change = data.image1.elo - current_winner_elo;
            loser_elo_change = data.image2.elo - current_loser_elo;

            console.log(winner_elo_change);
            console.log(loser_elo_change);

            show_result = true;

            await new Promise(f => setTimeout(f, 1000));

            show_result = false;

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

    async function handle_selection (n: number) {
        selected = n;

        if (n == 1) {
            winner_at_left = true;
            submit_vote(Image1, Image2)
        } else {
            winner_at_left = false;
            submit_vote(Image2, Image1)
        }

        await new Promise(f => setTimeout(f, 1000));
        await fetchRandomImages();
    }

    function handleKeyDown(e: KeyboardEvent) {
		 switch(e.key) {
			 case 'ArrowLeft':
                handle_selection(1);
				break;
			 case 'ArrowRight':
                handle_selection(2);
				break;
		 }
         console.log("KEYPRESSED");
	}
    
    // Load initial images
    onMount(() => {
        fetchRandomImages();
        window.addEventListener('keydown', handleKeyDown);
    });
</script>

<style>
* {
    padding: 0;
    margin: 0;
}


main {
    font-family: sans-serif;
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    align-items: center;
    height: 100vh;
    background-color: black;
    /* background: linear-gradient(#CAFFD0, #C9E4E7, #CA3CFF, #CA3CFF); */
}
h1 {
    position: absolute;
    top: 50%;
    transform: translateY(-50%);
    z-index: 3;
    color: black;
    font-weight: 800;
    font-size: 3rem;
    font-style: italic;
    text-shadow: 1px 0 #fff, -1px 0 #fff, 0 1px #fff, 0 -1px #fff,
            1px 1px #fff, -1px -1px #fff, 1px -1px #fff, -1px 1px #fff;
}
container {
    display: flex;
    flex-wrap: wrap;
    /* justify-content: space-around; */
    /* gap: 30px; */
    width: 100%;
    /* border: solid 2px blue; */
}

.layer {
    height: 100vh;
    width: 100vw;

    background-color: black;
    z-index: 4;
}


@media (max-width: 640px) {
    container {
        justify-content: start;
        gap: 0;
    }
    h1 {
        position: absolute;
        top: 50vh;
        transform: translateY(-50%);
        z-index: 3;
        color: black;
        font-weight: 800;
        font-size: 2rem;
        text-shadow: 1px 0 #fff, -1px 0 #fff, 0 1px #fff, 0 -1px #fff,
               1px 1px #fff, -1px -1px #fff, 1px -1px #fff, -1px 1px #fff;
    }
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

    {#if show_result}
        <div class="layer"> </div>
        {#if winner_at_left}
            <Pop {winner_elo_change} {loser_elo_change} winner_at_left={true} />
        {:else}
            <Pop {winner_elo_change} {loser_elo_change} winner_at_left={false} />
        {/if}
    {:else if loading}
        <div class="loading">Loading images...</div>
    {:else}
        <container>
            <Image image={Image1} handler={handle_selection} side="left" />
            <Image image={Image2} handler={handle_selection} side="right" />
        </container>
    {/if}
</main>