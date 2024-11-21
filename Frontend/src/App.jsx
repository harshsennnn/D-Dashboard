import React, { useEffect, useState } from 'react';

function App() {
  const [metrics, setMetrics] = useState(null);

  useEffect(() => {
    const fetchMetrics = async () => {
      const response = await fetch('<http://localhost:8080/api/stats>');
      const data = await response.json();
      setMetrics(data);
    };

    fetchMetrics();
    const interval = setInterval(fetchMetrics, 5000); // Poll every 5 seconds
    return () => clearInterval(interval);
  }, []);

  return (
    <div className="App">
      <h1>System Metrics</h1>
      {metrics ? (
        <div>
          <p>CPU Usage: {metrics.cpu_usage}%</p>
          <p>Memory Usage: {metrics.memory_usage}%</p>
          <p>Disk Usage: {metrics.disk_usage}%</p>
          <p>Network Usage: {metrics.network_usage}</p>
        </div>
      ) : (
        <p>Loading...</p>
      )}
    </div>
  );
}

export default App;
