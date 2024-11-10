import { setTimeout } from 'timers/promises';

async function simulateServiceLatency(min: number, max: number): Promise<number> {
    const latency = Math.floor(Math.random() * (max - min + 1)) + min;

    await setTimeout(latency);

    return latency;
}
export const lambdaHandler = async () => {
    const response = {
        statusCode: 200,
        body: JSON.stringify({
            message: `O serviço demorou ${await simulateServiceLatency(600, 1000)} ms para responder`,
        }),
    };

    return response;
};
