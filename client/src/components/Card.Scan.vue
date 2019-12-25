<template>
<section class="section">
    <div class="columns">
        <div class="column is-half is-offset-one-quarter">
            <div class="card">
            <header class="card-header">
                <p class="card-header-title">Report for {{ email.email }}</p>
                <!-- <span class="icon">
                    <i class="fas fa-angle-down" aria-hidden="true"></i>
                </span> -->
            </header>
            <div class="card-content">
                <div class="content">
                    Email <strong><u>{{ email.email }}</u></strong> has been reported <strong>{{ email.no_of_occurrences }}</strong> times for sending fraudulent emails.
                <br>
                First reported at <strong><time>{{ email.date_created }}</time></strong>
                <br>
                Last reported at <strong><time>{{ email.updated_at }}</time></strong>
                </div>
            </div>
            <footer class="card-footer">
                <a href="#" class="card-footer-item">+1 Report</a>
            </footer>
        </div>
        </div>
    </div>
</section>
</template>

<script>
import { mapGetters } from 'vuex';

export default {
    name: 'CardScan',
    computed: {
        ...mapGetters({ email: 'getEmail'})
    }
}
</script>