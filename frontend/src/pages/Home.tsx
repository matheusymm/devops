import { useState } from 'react';
import Stack from '@mui/material/Stack';
import IconButton from '@mui/material/IconButton';
import SentimentDissatisfiedIcon from '@mui/icons-material/SentimentDissatisfied';
import SentimentNeutralIcon from '@mui/icons-material/SentimentNeutral';
import SentimentSatisfiedIcon from '@mui/icons-material/SentimentSatisfied';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button';
import Header from "../components/Header";
import useMood from '../hooks/useMood';

const Home = () => {
  const token = localStorage.getItem("token");
  const userId = localStorage.getItem("userId");
  const [mood, setMood] = useState<number | null>(null);
  const [description, setDescription] = useState<string>("");
  const { createMood } = useMood();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (mood !== null && userId) {
      await createMood(userId, mood, description);
      setMood(null);
      setDescription("");
    }
  };

  return (
    <div className="flex flex-col bg-slate-300 w-screen h-screen">
      <Header />
      {token ? (
        <div className="flex w-full h-full">
          <form action="" className='flex flex-col justify-center items-center w-full h-full gap-5'>
            <h1 className="text-2xl mb-4">Como está se sentindo hoje?</h1>
            <Stack direction="row" spacing={1}>
              <IconButton aria-label="sad" title='Triste' color={mood === 1 ? "primary" : "default"} onClick={() => setMood(1)}>
                <SentimentDissatisfiedIcon fontSize='large' />
              </IconButton>
              <IconButton aria-label="neutral" title='Neutro' color={mood === 2 ? "primary" : "default"} onClick={() => setMood(2)}>
                <SentimentNeutralIcon fontSize='large' />
              </IconButton>
              <IconButton aria-label="happpy" title='Feliz' color={mood === 3 ? "primary" : "default"} onClick={() => setMood(3)}>
                <SentimentSatisfiedIcon fontSize='large' />
              </IconButton>
            </Stack>
            <TextField
              id="outlined-multiline-static"
              label="Descrição"
              multiline
              rows={4}
              value={description}
              onChange={(e) => setDescription(e.target.value)}
              className='w-1/2 my-4'
            />
            <Button variant="contained" color="primary" type="submit" size='large' onClick={handleSubmit}>
              Enviar
            </Button>
          </form>
        </div>
      ) : (
        <div className="flex flex-col items-center w-full h-full ">
          <h1 className="text-3xl font-bold text-center mt-10">Bem-vind ao Mood Tracker</h1>
          <p className="text-center mt-4">Faça login ou cadastre-se para ter acesso.</p>
        </div>
      )}
    </div>
  );
}

export default Home;