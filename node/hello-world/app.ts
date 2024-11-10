import { APIGatewayProxyEvent, APIGatewayProxyResult } from 'aws-lambda';
import crypto from 'crypto';

function encryptAES(key: string, text: string) {
    const iv = crypto.randomBytes(16);
    const cipher = crypto.createCipheriv('aes-256-cbc', Buffer.from(key), iv);
    let encrypted = cipher.update(text);
    encrypted = Buffer.concat([encrypted, cipher.final()]);
    return iv.toString('hex') + ':' + encrypted.toString('hex');
}

export const lambdaHandler = async (event: APIGatewayProxyEvent): Promise<APIGatewayProxyResult> => {
    const key = 'thisis32bitlongpassphraseimusing';
    const plaintext = 'Hello, AWS Lambda!';
    const encryptedResults = [];

    const timeStarted = Date.now();

    for (let i = 0; i < 10000; i++) {
        const encrypted = encryptAES(key, plaintext);
        encryptedResults.push(encrypted);
    }

    const timeElapsed = Date.now() - timeStarted;

    const response = {
        statusCode: 200,
        body: JSON.stringify({
            message: `O tempo de execução foi de: ${timeElapsed}ms, resultados criptografados: ${encryptedResults.length}`,
        }),
    };

    return response;
};
