"use strict;"

import axios from 'axios';

const BASE_URL = `http://localhost:4000/api/v1.0`;

async function ping() {
    return await axios.get(`${BASE_URL}/ping`);
}


export default {
    ping
}
