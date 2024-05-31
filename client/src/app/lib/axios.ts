'use server';

import axios from 'axios';

export const axiosInstance = axios.create({
  baseURL: 'https://api.kupolog.com',
  headers: {
    'Content-Type': 'application/json',
  },
});

export default axiosInstance;
