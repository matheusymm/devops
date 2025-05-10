import { useEffect } from "react";
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemText from '@mui/material/ListItemText';
import ListItemAvatar from '@mui/material/ListItemAvatar';
import SentimentDissatisfiedIcon from '@mui/icons-material/SentimentDissatisfied';
import SentimentNeutralIcon from '@mui/icons-material/SentimentNeutral';
import SentimentSatisfiedIcon from '@mui/icons-material/SentimentSatisfied';
import Avatar from '@mui/material/Avatar';
import Header from "../components/Header";
import useMood from "../hooks/useMood";
import { formatDate } from "../utils/date";

const MoodList = () => {
  const userId = localStorage.getItem("userId");
  const { moods, getMoodByUserId } = useMood();

  const fetchMoods = async () => {
    if (userId) {
      await getMoodByUserId(userId);
    }
  };

  useEffect(() => {
    fetchMoods();
  }, []);

  console.log(moods);

  return (
    <div className="flex flex-col bg-slate-300 w-screen h-screen">
      <Header />
      <div className="flex w-full h-full justify-center">
        {moods && moods.length === 0 && (
          <div className="flex flex-col justify-center items-center w-full h-full">
            <h1 className="text-2xl font-bold">Nenhum humor encontrado</h1>
            <p className="text-lg">Adicione um humor para começar a acompanhar seu humor.</p>
          </div>
        )}
        <List sx={{ width: '100%', maxWidth: 240 }}>
          {moods && moods.map((mood) => (
            <ListItem key={mood.id}>
              <ListItemAvatar>
                <Avatar>
                  {mood.mood === 1 ? (
                    <SentimentDissatisfiedIcon color="error" />
                  ) : mood.mood === 2 ? (
                    <SentimentNeutralIcon color="info" />
                  ) : (
                    <SentimentSatisfiedIcon color="success" />
                  )}
                </Avatar>
              </ListItemAvatar>
              <ListItemText primary={mood.description} secondary={formatDate(mood.created_at)} />
            </ListItem>
          ))}
        </List>
      </div>
    </div>
  );
}

export default MoodList;