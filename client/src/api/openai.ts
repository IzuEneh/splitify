  
  const apiEndpoint = "https://api.openai.com/v1/chat/completions";

  async function requestAutoPlaylist(apiKey: string, prompt_content:string) {
        
        const padded_prompt="Using the following list of comma delimited songs, can you provide \
            me a subset of these songs (Max 10) based on a unique theme of your choosing? \
            Please label the theme at the beginning of your message with simply Theme: <Theme> \
            Do not base the theme on the names, artist or album but on the content of the songs"+prompt_content

        const prompt = {
            "model": "gpt-3.5-turbo",
            "messages": [{
                "role": "user", 
                "content": padded_prompt
            }],
            "temperature": 0.5,
          }
    
        try {
            const response = await fetch(apiEndpoint, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${apiKey}`
                },
                body: JSON.stringify(prompt)
            });
  
            if (!response.ok) {
                throw new Error(`Error: ${response.status}`);
            }
  
            return response

        } catch (error) {
            throw error;
        }
    }
  
  export { requestAutoPlaylist }