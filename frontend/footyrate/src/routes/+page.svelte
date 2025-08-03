<script lang="ts">
    import IMG1 from '../images/1.jpg';
    import IMG2 from '../images/2.jpg';

</script>
<main>
    <div class="rate">
        <a href="/compare">
            <p id="rate-txt">Compare</p>
            <img src={IMG1} alt="compare">
        </a>
    </div>

    <div class="standings">
        <a href="/standings">
            <p id="standings-txt">Standings</p>
            <img src={IMG2} alt="standings">
        </a>
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
        overflow-y: hidden;

        background-color: black;
        text-shadow: 1px 0 black, -1px 0 black, 0 1px black, 0 -1px black,
               1px 1px black, -1px -1px black, 1px -1px black, -1px 1px black;

        color: bisque;
    }

    a:visited {
        color: red;
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

    #rate-txt {
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

    #standings-txt {
        position: absolute;
        top: 50%;
        left: 75%;
        transform: translateX(-50%);

        font-weight: 800;
        font-size: 3rem;
        font-style: italic;
    }

    a {
        display: block;
        height: 100%;
        width: 100%;
        text-decoration: none;
    }

    @media (max-width: 640px) {
        main {
            flex-direction: column;
            height: 100vh;
            overflow-y: hidden;
        }
        .rate {
            width: 100%;
            height: 50vh;
        }
        .standings {
            width: 100%;
            height: 50vh;
        }
        #rate-txt {
            top: 25%;
            left: 50%;
            transform: translateX(-50%);
        }
        #standings-txt {
            top: 75%;
            left: 50%;
            transform: translateX(-50%);
        }
    }
</style>
