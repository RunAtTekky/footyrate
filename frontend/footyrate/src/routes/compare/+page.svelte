<script>
    import Img1 from '../../images/1.webp'
    import Img2 from '../../images/2.webp'

    const IMAGE_WIDTH = '300px';
</script>

<style>
    * {
        padding: 0;
        margin: 0;
    }
    main {
        display: flex;
        justify-content: center;
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
</style>


<h1>Hola</h1>
<main>
    <container>
        <button>
            <img src={Img1} width={IMAGE_WIDTH} alt="Are you doing your part?">
        </button>
        <button>
            <img src={Img2} width={IMAGE_WIDTH} alt="Become spiderman or FROGGO">
        </button>
    </container>
</main>