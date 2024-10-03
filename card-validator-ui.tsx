import React, { useState } from 'react';

const CardValidator = () => {
    const [cardNumber, setCardNumber] = useState('');
    const [isValid, setIsValid] = useState<boolean | null>(null);
    const [error, setError] = useState<string | null>(null);

    const validateCard = async () => {
        setError(null);
        try {
            const response = await fetch('http://localhost:8080/validate-card', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ card_number: cardNumber }),
            });

            if (!response.ok) {
                throw new Error('Failed to validate card');
            }

            const data = await response.json();
            setIsValid(data.is_valid);
        } catch (err) {
            setError(err.message);
        }
    };

    return (
        <div>
            <h1>Credit Card Validator</h1>
            <input
                type="text"
                value={cardNumber}
                onChange={(e) => setCardNumber(e.target.value)}
                placeholder="Enter card number"
            />
            <button onClick={validateCard}>Validate</button>

            {isValid !== null && (
                <p>{isValid ? 'Card is valid!' : 'Card is invalid!'}</p>
            )}

            {error && <p style={{ color: 'red' }}>{error}</p>}
        </div>
    );
};

export default CardValidator;
