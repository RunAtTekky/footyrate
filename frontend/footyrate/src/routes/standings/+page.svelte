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

<main>
  <h1>Standings</h1>

  <table class="zigzag">
      <thead>

      <tr>
          <th class="header">NAME</th>
          <th class="header">K_FACTOR</th>
          <th class="header">ROUNDS</th>
          <th class="header">ELO</th>
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

</main>


<style>
main {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

h1 {
  color:#c00;
  font-family:sans-serif;
  font-size:2em;
  margin-bottom:0;
}

table {  
  font-family:sans-serif;
  th, td {
    padding:.25em .5em;
    text-align:left;
    &:nth-child(2) {
      text-align:right;
    }
  }
  td {
    background-color:#eee;    
  }
  th {
    background-color:#009;
    color:#fff;
  }
}

.zigzag {
  border-collapse:separate;
  border-spacing:.25em 1em;
  tbody tr:nth-child(odd) {
    transform:rotate(2deg);
  }
  thead tr,
  tbody tr:nth-child(even) {
    transform:rotate(-2deg);
  }
}
</style>