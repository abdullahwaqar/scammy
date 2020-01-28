<template>
<section class="section">
    <div class="columns">
        <div class="column is-half is-offset-one-quarter">
            <!-- If no data -->
            <div v-if="!email.id" class="card">
                <div class="card-content">
                    <center>
                        <a href="/" class="fas fa-exclamation fa-5x"></a>
                    </center>
                </div>
            </div>
            <div v-else class="card">
            <header class="card-header">
                <p class="card-header-title">Report for {{ email.email }}</p>
                <!-- <span class="icon">
                    <i class="fas fa-angle-down" aria-hidden="true"></i>
                </span> -->
            </header>
            <div class="card-content">
                <div class="tabs">
                <ul>
                    <li @click="tab = 'report'" :class="[tab === 'report' ? 'is-active' : '']"><a>Report</a></li>
                    <li @click="tab = 'header'" :class="[tab === 'header' ? 'is-active' : '']"><a>Sample Email Header</a></li>
                </ul>
                </div>
                <!-- Report tab -->
                <div v-if="tab === 'report'">
                    <div class="content">
                        Email <strong><u>{{ email.email }}</u></strong> has been reported <strong>{{ email.no_of_occurrences }}</strong> times for sending fraudulent emails.
                    <br>
                    First reported at <strong><time>{{ email.date_created }}</time></strong>
                    <br>
                    Last reported at <strong><time>{{ email.updated_at }}</time></strong>
                    </div>
                </div>
                <!-- Header Tab -->
                <div v-if="tab === 'header'">
                    <div class="control">
                    <textarea v-model="email.email_header" class="textarea has-fixed-size" placeholder="Fixed size textarea" readonly></textarea>
                    </div>
                </div>
            </div>
            <footer class="card-footer">
                <a v-if="flag" @click="increment" class="card-footer-item">+1 Report</a>
            </footer>
        </div>
        </div>
    </div>
</section>
</template>

<script>
import { mapGetters, mapActions } from 'vuex';

export default {
    name: 'CardScan',
    data: function() {
        return {
            tab: 'report',
            flag: true
        }
    },
    computed: {
        ...mapGetters({ email: 'getEmail' })
    },
    methods: {
        ...mapActions({ incrementEmail: 'incrementReportCount' }),
        increment() {
            this.incrementEmail();
            this.flag = false;
        }
    }
}
</script>