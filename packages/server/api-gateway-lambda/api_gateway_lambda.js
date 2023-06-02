const axios = require('axios');
const { Client } = require('pg');

export const handler = async(event) => {
    try {
        // Retrieve RDS connection configuration from environment variables
        const rdsConfig = {
            host: process.env.RDS_HOST,
            port: process.env.RDS_PORT,
            user: process.env.RDS_USERNAME,
            password: process.env.RDS_PASSWORD,
            database: process.env.RDS_DATABASE
        };

        // Read the X-OPENLINE-TENANT-KEY header from the event
        const tenantKey = event.headers['X-OPENLINE-TENANT-KEY'];

        // Create PostgreSQL client
        const client = new Client(rdsConfig);

        // Connect to the PostgreSQL database
        await client.connect();

        // Query the tenantkeys table based on the tenantKey
        const query = `SELECT tenant, key, active FROM tenant_keys WHERE tenant = $1`;
        const result = await client.query(query, [tenantKey]);
        const rows = result.rows;

        // Disconnect from the PostgreSQL database
        await client.end();

        if (rows.length > 0) {
            const tenant = rows[0].tenant;

            // Prepare the request to the targetAPI
            const targetAPIUrl = process.env.TARGET_API_URL;
            const headers = {
                'X-OPENLINE-TENANT': tenant,
                'X-OPENLINE-KEY': tenantKey
            };

            // Make a POST request to the targetAPI
            const response = await axios.post(targetAPIUrl, event.body, { headers });

            // Log the response from the targetAPI
            console.log('Response from targetAPI:', response.data);

            return {
                statusCode: 200,
                body: JSON.stringify(response.data)
            };
        } else {
            return {
                statusCode: 404,
                body: JSON.stringify({ error: 'Tenant key not found' })
            };
        }
    } catch (error) {
        console.error('Error:', error);

        return {
            statusCode: 500,
            body: JSON.stringify({ error: 'Internal Server Error' })
        };
    }
};