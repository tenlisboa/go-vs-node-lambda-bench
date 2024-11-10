import { APIGatewayProxyEvent, APIGatewayProxyResult } from 'aws-lambda';

function isPrime(n: number) {
    if (n <= 1) return false;
    for (let i = 2; i <= Math.sqrt(n); i++) {
        if (n % i === 0) return false;
    }
    return true;
}

function calculatePrimes(limit: number) {
    const primes = [];
    for (let i = 2; i <= limit; i++) {
        if (isPrime(i)) {
            primes.push(i);
        }
    }
    return primes;
}
export const lambdaHandler = async (event: APIGatewayProxyEvent): Promise<APIGatewayProxyResult> => {
    const limit = 1000000;

    const primes = calculatePrimes(limit);
    console.log();

    const response = {
        statusCode: 200,
        body: JSON.stringify({
            message: `Números primos até ${limit}: ${primes.length}`,
        }),
    };

    return response;
};
