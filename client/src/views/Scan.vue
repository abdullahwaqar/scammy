<template>
    <section>
        <header>
            <navbar></navbar>
        </header>
        <br>
        <main>
            <card></card>
        </main>
        <footer class="footer-bt"><cs-footer></cs-footer></footer>
    </section>
</template>

<script>
import Navbar from '@/components/Navbar.vue';
import CardScan from '@/components/Card.Scan.vue';
import Footer from '@/components/Footer.vue';

export default {
    name: 'Scan',
    components: {
        'navbar': Navbar,
        'card': CardScan,
        'cs-footer': Footer
    }
}
</script>

<style scoped>
.footer-bt {
    position: absolute;
    bottom: 0;
    height: 20px;
    width: 100%;
}
</style>