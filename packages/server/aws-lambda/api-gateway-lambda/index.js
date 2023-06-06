const https = require('https');
const { Client } = require('pg');

exports.handler = async (event) => {
    try {
        console.log('Event:', JSON.stringify(event));
        if (!event.body) {
            return {
                statusCode: 400,
                body: JSON.stringify({ error: 'Invalid request body' }),
            };
        }
        // Retrieve RDS connection configuration from environment variables
        const rdsConfig = {
            host: process.env.RDS_HOST,
            port: process.env.RDS_PORT,
            user: process.env.RDS_USERNAME,
            password: process.env.RDS_PASSWORD,
            database: process.env.RDS_DATABASE,
        };

        // Read the X-OPENLINE-TENANT-KEY header from the event
        const tenantKey = event.headers['x-openline-tenant-key'];

        // Create PostgreSQL client
        const client = new Client(rdsConfig);

        // Connect to the PostgreSQL database
        console.log('Connecting to PostgreSQL database...');
        await client.connect();

        // Query the tenant_keys table based on the tenantKey
        const query = `SELECT tenant, key, active FROM tenant_keys WHERE key = $1`;
        const result = await client.query(query, [tenantKey]);
        const rows = result.rows;

        // Disconnect from the PostgreSQL database
        await client.end();

        if (rows.length > 0) {
            const tenant = rows[0].tenant;

            // Prepare the request to the targetAPI
            const targetAPIUrl = process.env.TARGET_API_URL;

            const headers = {
                'X-openline-TENANT': tenant,
                'X-openline-API-KEY': process.env.X_Openline_API_KEY,
            };

            // Make a POST request to the targetAPI
            console.log('Calling target API...' + targetAPIUrl);

            const options = {
                method: 'POST',
                headers: {
                    ...headers,
                    'Content-Type': 'application/json',
                    'Content-Length': Buffer.byteLength(event.body),
                },
            };

            const response = await new Promise((resolve, reject) => {
                const req = https.request(targetAPIUrl, options, (res) => {
                    let data = '';

                    res.on('data', (chunk) => {
                        data += chunk;
                    });

                    res.on('end', () => {
                        resolve({
                            statusCode: res.statusCode,
                            body: data,
                        });
                    });
                });

                req.on('error', (error) => {
                    reject(error);
                });

                req.write(event.body);
                req.end();
            });

            // Log the response from the targetAPI
            console.log('Response from targetAPI:', response.body);

            return {
                statusCode: response.statusCode,
                body: JSON.stringify(response.body),
            };
        } else {
            return {
                statusCode: 404,
                body: JSON.stringify({ error: 'Tenant key not found' }),
            };
        }
    } catch (error) {
        console.error('Error:', error);

        return {
            statusCode: 500,
            body: JSON.stringify({ error: 'Internal Server Error' }),
        };
    }
};
