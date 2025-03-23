<script>
    import Img1 from '../../images/1.webp'
    import Img2 from '../../images/2.webp'

    let selected = $state(0);

    const IMAGE_WIDTH = '300px';
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
        width: 80%;
        background-color: #C9E4E7;
        border-radius: 10px;
    }

    button {
        margin: 20px;
        border: none;
    }

    .result {
        padding-top: 20px;
    }
</style>


<main>
    <h1>Hola</h1>
    <container>
        <button onclick={() => {selected = 1;}}>
            <img src={Img1} width={IMAGE_WIDTH} alt="Are you doing your part?">
        </button>
        <button onclick={() => {selected = 2;}}>
            <img src={Img2} width={IMAGE_WIDTH} alt="Become spiderman or FROGGO">
        </button>
    </container>

    {#if selected != 0}
    <p class="result">Image {selected} selected</p>
    {/if}

</main>