<template>
<section>
    <main>
        <navbar></navbar>
    </main>
    <footer>
        <footer-main></footer-main>
    </footer>
</section>
</template>

<script>
import MainPage from '@/components/Mainpage.vue';
import Footer from '@/components/Footer.vue';

export default {
    name: 'home',
    components: {
        'navbar': MainPage,
        'footer-main': Footer
    }
}
</script>
