const express = require('express');
require('dotenv').config();

const statisticsRoutes = require('./routes/statisticsRoutes');

const app = express();
const port = process.env.PORT || 3000;

app.use(express.json());

app.get('/', (req, res) => {
  res.json({message: 'Node API is running'});
});

app.use('/', statisticsRoutes);

app.listen(port, () => {
  console.log(`Output Server is running on port ${port}`);
});