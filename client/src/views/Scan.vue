<template>
    <section>
        <header>
            <navbar></navbar>
        </header>
        <br>
        <main>
            <card></card>
        </main>
        <footer></footer>
    </section>
</template>

<script>
import Navbar from '@/components/Navbar.vue';
import CardScan from '@/components/Card.Scan.vue';

export default {
    name: 'Scan',
    components: {
        'navbar': Navbar,
        'card': CardScan
    }
}
</script>