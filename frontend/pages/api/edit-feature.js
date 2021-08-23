import axios from 'axios';
import Cors from 'cors';
import initMiddleware from '../../lib/init-middleware';

// Initialize the cors middleware
const cors = initMiddleware(
    // You can read more about the available options here: https://github.com/expressjs/cors#configuration-options
    Cors({
    // Only allow requests with GET, POST and OPTIONS
        methods: ['GET', 'POST', 'OPTIONS'],
    }),
);

export default async function handler(req, res) {
    // Run cors
    await cors(req, res);

    // Rest of the API logic
    const url = 'http://localhost:8080/api/v1/feature';

    try {
        const { data } = await axios.put(url, req.body, {
            headers: { 'Content-Type': 'application/json' },
        });

        res.json({ data });
    } catch (error) {
        console.log(error.response.data);
    }
}
