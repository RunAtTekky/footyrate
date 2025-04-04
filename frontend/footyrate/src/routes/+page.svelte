<script lang="ts">
    import IMG1 from '../images/1.jpg';
    import IMG2 from '../images/2.jpg';

</script>
<main>
    <div class="rate">
        <a href="/compare">
            <img src={IMG1} alt="compare">
        </a>
        <p>Compare</p>
    </div>

    <div class="standings">
        <a href="/standings">
            <img src={IMG2} alt="standings page">
        </a>
        <p>Standings</p>
    </div>
</main>



<style>
    main {
        font-family: sans-serif;
        height: 100vh;
        display: flex;
        flex-direction: column;
        justify-content: flex-start;
        align-items: center;

        background-color: bisque;

        color: bisque;
    }

    img {
        height: 100%;
        width: 100%;
        object-fit: cover;
        /* filter: blur(1px); */
    }

    .rate {
        height: 50vh;
        width: 100%;
    }

    .rate > p {
        position: absolute;
        top: 20%;
        left: 50%;
        transform: translateX(-50%);

        font-weight: 800;
        font-size: 3rem;
        font-style: italic;
    }

    .standings {
        height: 50vh;
        width: 100%;
    }

    .standings > p {
        position: absolute;
        top: 70%;
        left: 50%;
        transform: translateX(-50%);

        font-weight: 800;
        font-size: 3rem;
        font-style: italic;
    }

    @media (max-width: 450px) {
        .rate > p {
            top: 25%;
        }
        .standings > p {
            top: 65%;
        }
    }
</style>