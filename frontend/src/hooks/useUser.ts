import { useState } from 'react';
import { useNavigate } from 'react-router';
import api from '../services/api';

const useUser = () => {
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState(null);
    const navigate = useNavigate();

    const login = async (email: string, password: string) => {
        setLoading(true);
        const hashedPassword = btoa(password);
        await api.post('/users/login', { email, password: hashedPassword })
            .then((response) => {
                const data = response.data;
                localStorage.setItem('token', data.token);
                localStorage.setItem('userId', data.user_id);
                navigate('/');
            }
            )
            .catch((err) => {
                setError(err.response.data.message);
            }
            )
            .finally(() => {
                setLoading(false);
            }
            );
    };

    const signup = async (name: string, email: string, password: string) => {
        setLoading(true);
        const hashedPassword = btoa(password);
        await api.post('/users/', { name, email, password: hashedPassword })
            .then((response) => {
                const data = response.data;
                localStorage.setItem('token', data.token);
                localStorage.setItem('userId', data.user_id);
                navigate('/');
            }
            )
            .catch((err) => {
                setError(err.response.data.message);
            }
            )
            .finally(() => {
                setLoading(false);
            }
            );
    };

    const logout = () => {
        localStorage.removeItem('token');
        localStorage.removeItem('userId');
        navigate('/');
    };

    return { loading, error, login, signup, logout };
}

export default useUser;