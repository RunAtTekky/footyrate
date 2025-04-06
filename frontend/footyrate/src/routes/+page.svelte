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
        /* flex-direction: column; */
        justify-content: flex-start;
        align-items: center;

        background-color: black;
        text-shadow: 1px 0 black, -1px 0 black, 0 1px black, 0 -1px black,
               1px 1px black, -1px -1px black, 1px -1px black, -1px 1px black;

        color: bisque;
    }

    img {
        height: 100%;
        width: 100%;
        object-fit: cover;
        /* filter: blur(1px); */
    }

    .rate {
        width: 50vw;
        height: 100vh;
    }

    .rate > p {
        position: absolute;
        top: 50%;
        left: 25%;
        transform: translateX(-50%);

        font-weight: 800;
        font-size: 3rem;
        font-style: italic;
    }

    .standings {
        width: 50vw;
        height: 100vh;
    }

    .standings > p {
        position: absolute;
        top: 50%;
        left: 75%;
        transform: translateX(-50%);

        font-weight: 800;
        font-size: 3rem;
        font-style: italic;
    }

    @media (max-width: 640px) {
        main {
            flex-direction: column;
        }
        .rate {
            width: 100%;
        }
        .standings {
            width: 100%;
        }
        .rate > p {
            top: 25%;
            left: 50%;
            transform: translateX(-50%);
        }
        .standings > p {
            top: 75%;
            left: 50%;
            transform: translateX(-50%);
        }
    }
</style>