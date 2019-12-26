"use strict;"

import axios from 'axios';

const BASE_URL = `http://localhost:4000/api/v1.0`;

async function ping() {
    return await axios.get(`${BASE_URL}/ping`);
}

async function getEmailDEtails(email) {
    try {
        const response = await axios.post(`${BASE_URL}/scan`, { email: email });
        return response.data;
    } catch (error) {
        return { error: error };
    }
}

async function incrementEmailReport(id) {
    try {
        const response = await axios.patch(`${BASE_URL}/${id}`, { vote: 1 });
        return response.data;
    } catch (error) {
        return { error: error };
    }
}

async function postNewScamEmail(email, emailHeader) {
    try {
        const response = await axios.post(`${BASE_URL}/newscam`, { email: email, email_header: emailHeader });
        return response.data;
    } catch (error) {
        console.log(error)
        return { error: error };
    }
}

export default {
    ping,
    getEmailDEtails,
    incrementEmailReport,
    postNewScamEmail
}
