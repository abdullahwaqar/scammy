<template>
    <section>
        <header>
            <navbar></navbar>
        </header>
        <br>
        <main>
            <card></card>
            <div class="columns">
                <div class="column is-half is-offset-one-quarter">
                <article class="message is-dark">
                <div class="message-body">
                    If you have recieved an email from this address we suggest you to not engage with it. Please consider hitting <strong><u>+1 Report</u></strong> to help others.
                    <br>
                    <br>
                    If this email is not in our database and you think this email is used to scam people then please report it <strong><a href="/report">here</a></strong>.
                    <br>
                    <br>
                    <strong>Tips:</strong>
                    <ul>
                        <li>1. Don't get into deal with wired money.</li>
                        <li>3. Never get into sweepstakes/foreign lotteries.</li>
                        <li>2. Never reply to messages asking you for financial/personal information.</li>
                    </ul>
                </div>
                </article>
                </div>
            </div>
        </main>
        <footer class="footer-bt"><cs-footer></cs-footer></footer>
    </section>
</template>

<script>
import Navbar from '@/components/Navbar.vue';
import CardScan from '@/components/Card.Scan.vue';
import Footer from '@/components/Footer.vue';

export default {
    name: 'Report',
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