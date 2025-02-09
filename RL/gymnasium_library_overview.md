# Gymnasium Library - Overview

## What is Gymnasium?
**Gymnasium** is an open-source Python library used to create and train reinforcement learning (RL) agents. It provides a standard API for RL environments, making it easy to develop and benchmark RL algorithms.

Gymnasium is a maintained fork of the original **OpenAI Gym**, offering better long-term support and additional features.

---

## Features of Gymnasium
1. **Standardized API**: Provides a consistent interface to interact with various RL environments.
2. **Diverse Environments**: Supports multiple RL problems like classic control, Atari games, robotics, and more.
3. **Support for Vectorized Environments**: Allows running multiple instances of an environment for efficient training.
4. **Automatic Reset Handling**: Helps manage episode termination efficiently.
5. **Extended Support & Maintenance**: Aims for long-term stability with continued improvements.

---

## Installation
To install Gymnasium, use:

```bash
pip install gymnasium
```

For additional environments like `Atari` or `Mujoco`:

```bash
pip install gymnasium[atari]
pip install gymnasium[mujoco]
```

---

## Basic Example: Using Gymnasium in Python

Below is a simple example of how to use Gymnasium with the **CartPole** environment.

```python
import gymnasium as gym

# Create the CartPole environment
env = gym.make("CartPole-v1", render_mode="human")

# Reset the environment before starting
obs, info = env.reset()

for _ in range(1000):  # Run for 1000 steps
    env.render()  # Render the environment visually
    
    action = env.action_space.sample()  # Choose a random action
    
    # Step the environment forward by taking the chosen action
    obs, reward, terminated, truncated, info = env.step(action)

    # Check if the episode is over (either by termination or truncation)
    if terminated or truncated:
        obs, info = env.reset()  # Restart the environment

env.close()  # Close the environment when done
```

---

## Explanation of Code

1. **`import gymnasium as gym`**  
   - Imports the Gymnasium library.

2. **`env = gym.make("CartPole-v1", render_mode="human")`**  
   - Creates an instance of the **CartPole-v1** environment.  
   - `render_mode="human"` enables visual rendering.

3. **`obs, info = env.reset()`**  
   - Resets the environment and gets the initial observation and metadata.

4. **Loop over 1000 steps:**
   - **`env.render()`**: Displays the environment.
   - **`action = env.action_space.sample()`**: Chooses a random action from the action space.
   - **`obs, reward, terminated, truncated, info = env.step(action)`**:  
     - `obs`: The new observation after the action.  
     - `reward`: The reward received.  
     - `terminated`: Whether the episode ended normally.  
     - `truncated`: Whether the episode was forcibly stopped (time limit).  
     - `info`: Extra debugging info.  
   - **If `terminated` or `truncated` is `True`**, reset the environment.

5. **`env.close()`**  
   - Closes the environment properly.

---

## Summary
- **Gymnasium** is a library for RL environment simulation.
- It provides an easy API to create, interact with, and train RL agents.
- Supports various domains, including games, robotics, and control tasks.

Would you like an example using a specific RL algorithm like Q-learning or PPO?

