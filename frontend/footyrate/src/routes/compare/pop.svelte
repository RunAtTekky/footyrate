<script lang="ts">
    const { winner_elo_change, loser_elo_change, winner_at_left} = $props();
</script>

{#if winner_at_left}
    <h2 class="elo_change left_side" style="z-index: 5; color:green">
        +{winner_elo_change}
    </h2>
    <h2 class="elo_change right_side" style="z-index: 5; color:red">
        {loser_elo_change}
    </h2>
{:else}
    <h2 class="elo_change right_side" style="z-index: 5; color:green">
        +{winner_elo_change}
    </h2>
    <h2 class="elo_change left_side" style="z-index: 5; color:red">
        {loser_elo_change}
    </h2>
{/if}

<style>

.elo_change {
    display: flex;
    z-index: 5;
    font-size: 3rem;
    font-weight: 800;
    position: absolute;
    animation: dropIn 0.6s cubic-bezier(.82,.13,.66,1.15) forwards, pulse 0.5s 0.6s ease-in-out;
}

.left_side {
    top: 50vh;
    left: 25vw;
    transform: translate(-50%, -50%);
}

.right_side {
    top: 50vh;
    left: 75vw;
    transform: translate(-50%, -50%);
}

@keyframes dropIn {
  0% {
    transform: translate(-50%, -50%) scale(4);
  }
  50% {
    transform: translate(-50%, -50%) scale(1.3);
  }
  70% {
    transform: translate(-50%, -50%) scale(1.2);
  }
  100% {
    transform: translate(-50%, -50%) scale(1);
  }
}

@keyframes pulse {
  0% {
    transform: translate(-50%, -50%) scale(1);
    text-shadow: 0 0 0px rgba(255, 255, 255, 0.8);
  }
  50% {
    transform: translate(-50%, -50%) scale(1.2);
    text-shadow: 0 0 20px rgba(255, 255, 255, 0.8);
  }
  100% {
    transform: translate(-50%, -50%) scale(1);
    text-shadow: 0 0 0px rgba(255, 255, 255, 0.8);
  }
}

@media (max-width: 640px) {
    .left_side {
        top: 25vh;
        left: 50vw;
    }
    .right_side {
        top: 75vh;
        left: 50vw;
    }
}

</style>