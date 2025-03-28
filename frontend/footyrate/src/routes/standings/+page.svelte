<script lang="ts">
	import { onMount } from "svelte";

    interface ImageData {
        id: number;
        url: string;
        elo: number;
        k_factor: number;
        rounds: number;
    }

    let Image_List: ImageData[];

    // const API_RESULT = "http://localhost:8080/api/images"
    const API_RESULT = "https://footyrate.onrender.com/api/images"

    function sort_standings() {
        Image_List.sort((a,b) => b.elo - a.elo);
    }

    async function get_standings() {
        try {
            const response = await fetch(API_RESULT);

            if (!response.ok) {
                throw new Error("Error getting standings");
            }

            const data = await response.json();
            
            Image_List = data;
        } catch (error) {
            console.error("Error fetching standings", error);
        }

        sort_standings();
    }

    onMount(() => {
        get_standings();
    });
</script>

<h2>Standings</h2>

<table>
    <thead>

    <tr>
        <th>NAME</th>
        <th>K_FACTOR</th>
        <th>ROUNDS</th>
        <th>ELO</th>
    </tr>
    </thead>

    <tbody>
    {#each Image_List as image}
    <tr>
            <td>{image.url}</td>
            <td>{image.k_factor}</td>
            <td>{image.rounds}</td>
            <td>{image.elo}</td>
    </tr>
    {/each}
    </tbody>
</table>