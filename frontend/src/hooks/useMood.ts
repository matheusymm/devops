import { useState } from 'react';
import { useNavigate } from 'react-router';
import api from '../services/api';
import { Mood } from '../types/mood';

const useMood = () => {
    const [moods, setMoods] = useState<Mood[] | null>(null);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState(null);
    const navigate = useNavigate();

    const createMood = async (userId: string, mood: number, description: string) => {
        setLoading(true);
        console.log(userId, mood, description);
        await api.post('/moods/', { user_id: userId, mood, description })
            .then((response) => {
                console.log(response.data);
                navigate('/moods');
            })
            .catch((err) => {
                setError(err.response.data.message);
            })
            .finally(() => {
                setLoading(false);
            });
    }

    const getMoodByUserId = async (userId: string) => {
        setLoading(true);
        await api.get(`/moods/user/${userId}`)
            .then((response) => {
                setMoods(response.data);
            })
            .catch((err) => {
                setError(err.response.data.message);
            })
            .finally(() => {
                setLoading(false);
            });
    };

    return { moods, loading, error, createMood, getMoodByUserId };
}

export default useMood;