<template>
<section class="section">
    <div class="columns">
        <div class="column is-half is-offset-one-quarter">
            <div v-if="emailRet.id" class="card">
                <div class="card-content">
                    <center>
                        <i class="fas fa-check-double fa-5x"></i>
                        <br>
                        <br>
                        Thank you for reporting. <u>{{ emailRet.email }}</u>
                    </center>
                </div>
        </div>
        <div v-if="!emailRet.id">
        <div class="field">
        <label class="label">Email *</label>
        <div class="control has-icons-left has-icons-right">
            <input v-model="email" :class="['input', !valid ? 'is-danger' : '']" type="email" placeholder="Email">
            <span class="icon is-small is-left">
            <i class="fas fa-envelope"></i>
            </span>
        </div>
        <p v-if="!valid" class="help is-danger">This email is invalid</p>
        </div>

        <div class="field">
        <label class="label">Email Headers (Optional)</label>
        <div class="control">
            <textarea v-model="emailHeader" class="textarea" placeholder="Email Header"></textarea>
        </div>
        </div>

        <div class="field is-grouped">
        <div class="control">
            <button @click="submit" class="button is-primary">Submit</button>
        </div>
        </div>
        </div>
        </div>
    </div>

</section>
</template>

<script>
import { mapActions, mapGetters } from 'vuex';

export default {
    name: 'ReportCard',
    data: function() {
        return {
            valid: true,
            email: '',
            emailHeader: ''
        }
    },
    computed: {
        ...mapGetters({ emailRet: 'getEmail'})
    },
    methods: {
        ...mapActions({ report: 'reportNewScam' }),
        validateEmail(email) {
            var re = /^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
            return re.test(email);
        },
        submit() {
            if (!this.validateEmail(this.email)) {
                this.valid = false;
                return;
            }
            const emailHeader = this.emailHeader.split('\n');
            this.report({
                email: this.email,
                emailHeader: emailHeader.filter(String).join()
            });
        }
    },
    watch: {
        email: function() {
            this.valid = true;
        }
    }
}
</script>