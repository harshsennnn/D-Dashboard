import React, { useEffect, useState } from 'react';

function App() {
  const [metrics, setMetrics] = useState(null);

  useEffect(() => {
    const fetchMetrics = async () => {
      const response = await fetch('<http://localhost:8080/api/stats>');
      const data = await response.json();
      setMetrics(data);
    };

    