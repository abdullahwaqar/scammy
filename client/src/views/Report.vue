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
                    If you have recieved an email from a scammer please contribute by submiting the email of the sender.
                    <br>
                    <br>
                    Or check an email if it is marked as scam in our database <strong><a href="/">here</a></strong>.
                    <br>
                    <br>
                    <u><a href="https://www.purpledogdesign.com/clients/knowledgebase/117/How-to-Copy-Email-Headers-for-troubleshooting.html">
                    What are Email headers? And where to find them.
                    </a></u>
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
import ReportCard from '@/components/Report.Card.vue';
import Footer from '@/components/Footer.vue';

export default {
    name: 'Report',
    components: {
        'navbar': Navbar,
        'card': ReportCard,
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