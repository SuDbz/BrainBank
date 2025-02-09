# LunarLander-v2



## Setting up the environment

```python
import gymnasium as gym

# Create the LunarLander-v3 environment
env = gym.make('LunarLander-v3')

# Get the shape of the observation space
state_shape = env.observation_space.shape  # Tuple representing the observation dimensions
state_size = env.observation_space.shape[0]  # Number of features in the state vector

# Get the number of possible actions
number_actions = env.action_space.n  # The total number of discrete actions available

# Print extracted information
print('State shape: ', state_shape)  # Example output: (8,)
print('State size: ', state_size)  # Example output: 8
print('Number of actions: ', number_actions)  # Example output: 4

```

## Initializing the hyperparameters

# Explanation of Key Hyperparameters

## 1. Learning Rate (`learning_rate = 5e-4`)
- **Definition**: Controls the step size at which the model updates during training.
- **Value**: `5e-4` (or `0.0005`)
- **Effect**: A smaller value results in slower learning but more stable updates, while a larger value can lead to faster learning but might cause instability.

## 2. Mini-batch Size (`minibatch_size = 100`)
- **Definition**: The number of samples processed before updating the model.
- **Value**: `100`
- **Effect**: Affects training stability and computational efficiency. A larger batch size smooths the gradient updates, while a smaller batch size introduces more noise.

## 3. Discount Factor (`discount_factor = 0.99`)
- **Definition**: Determines how much future rewards are valued compared to immediate rewards.
- **Value**: `0.99`
- **Effect**: A higher value (`≈1`) makes the agent focus on long-term rewards, while a lower value prioritizes short-term rewards.

## 4. Replay Buffer Size (`replay_buffer_size = int(1e5)`)
- **Definition**: The maximum number of past experiences stored for training.
- **Value**: `100,000` (`1e5`)
- **Effect**: Larger buffers provide more diverse experiences but require more memory, while smaller buffers may lead to less effective learning due to limited sample diversity.

## 5. Interpolation Parameter (`interpolation_parameter = 1e-3`)
- **Definition**: Used in soft updates of the target network in deep reinforcement learning.
- **Value**: `0.001` (`1e-3`)
- **Effect**: Controls the smoothness of updates to the target network. A lower value ensures more stable training by slowly incorporating new information.


# Experience Replay

# Explanation of ReplayMemory Class

## Overview
The `ReplayMemory` class is a data structure used in reinforcement learning to store past experiences and sample them randomly for training. This is essential for stabilizing learning and improving sample efficiency.

## Class Definition
### 1. **`__init__` Method**
```python
 def __init__(self, capacity):
    self.device = torch.device("cuda:0" if torch.cuda.is_available() else "cpu")  # Determines the device (GPU if available, else CPU)
    self.capacity = capacity  # Sets the maximum size of the replay buffer
    self.memory = []  # Initializes an empty list to store experiences
```
- **`self.device`**: Determines whether to use GPU (`cuda`) or CPU for computations, which speeds up training when using deep learning models.
- **`self.capacity`**: The maximum number of experiences that can be stored before old experiences start getting removed.
- **`self.memory`**: A list used to store experiences, which will grow dynamically up to the defined `capacity`.

### 2. **`push` Method**
```python
 def push(self, event):
    self.memory.append(event)  # Adds the new experience (event) to the memory
    if len(self.memory) > self.capacity:  # If memory exceeds capacity
      del self.memory[0]  # Remove the oldest experience to maintain buffer size
```
- **Purpose**: Adds new experiences to the memory while maintaining a fixed buffer size.
- **If the memory exceeds capacity**: The oldest experience is removed (`del self.memory[0]`), ensuring the buffer doesn't grow indefinitely.

### 3. **`sample` Method**
```python
 def sample(self, batch_size):
    experiences = random.sample(self.memory, k=batch_size)  # Randomly selects `batch_size` number of experiences from memory
    
    states = torch.from_numpy(np.vstack([e[0] for e in experiences if e is not None])).float().to(self.device)  # Extracts states and converts to tensor
    actions = torch.from_numpy(np.vstack([e[1] for e in experiences if e is not None])).long().to(self.device)  # Extracts actions taken in those states
    rewards = torch.from_numpy(np.vstack([e[2] for e in experiences if e is not None])).float().to(self.device)  # Extracts rewards received
    next_states = torch.from_numpy(np.vstack([e[3] for e in experiences if e is not None])).float().to(self.device)  # Extracts next states after taking actions
    dones = torch.from_numpy(np.vstack([e[4] for e in experiences if e is not None]).astype(np.uint8)).float().to(self.device)  # Converts done flags (True/False) to tensor
    
    return states, next_states, actions, rewards, dones  # Returns the sampled experiences as tensors
```
- **Purpose**: Randomly samples a batch of experiences from memory for training.
- **Extracts**:
  - `states`: Current state of the environment.
  - `actions`: Actions taken in those states.
  - `rewards`: Rewards received after performing actions.
  - `next_states`: Next states observed after taking actions.
  - `dones`: Boolean flags indicating whether an episode has ended (converted to integer 0 or 1 for computation).
- **Data Processing**:
  - Uses `np.vstack` to stack experiences into arrays.
  - Converts NumPy arrays to PyTorch tensors (`torch.from_numpy`).
  - Moves tensors to the appropriate device (CPU/GPU) using `.to(self.device)`, optimizing training speed.

## Summary
- **Replay memory is crucial** in reinforcement learning as it helps break correlations in sequential data.
- **It enables batch updates**, improving learning efficiency and stability.
- **Random sampling** ensures that learning is unbiased and covers diverse experiences, improving generalization.
- **Stores and samples tuples** containing (state, action, reward, next state, done) to train deep Q-networks (DQN) or other RL models.



# Explanation of the Agent Class

## Overview
The `Agent` class is responsible for interacting with the environment, storing experiences in memory, selecting actions, and learning from past experiences to improve decision-making. It implements the Deep Q-Learning algorithm using a neural network.

## Class Definition

### 1. **`__init__` Method**
```python
  def __init__(self, state_size, action_size):
    self.device = torch.device("cuda:0" if torch.cuda.is_available() else "cpu")  # Selects GPU if available, otherwise CPU
    self.state_size = state_size  # Number of dimensions in the state space
    self.action_size = action_size  # Number of possible actions
    self.local_qnetwork = Network(state_size, action_size).to(self.device)  # Main network used for selecting actions
    self.target_qnetwork = Network(state_size, action_size).to(self.device)  # Target network for stable learning
    self.optimizer = optim.Adam(self.local_qnetwork.parameters(), lr=learning_rate)  # Adam optimizer for updating network weights
    self.memory = ReplayMemory(replay_buffer_size)  # Experience replay buffer
    self.t_step = 0  # Step counter for controlling the frequency of learning updates
```
- **Two networks (`local_qnetwork` & `target_qnetwork`)**: Used in Deep Q-Learning to stabilize training.
- **Replay Memory (`ReplayMemory`)**: Stores experiences for experience replay.
- **Optimizer (`Adam`)**: Updates model weights using the loss function.
- **Device (`cuda` or `cpu`)**: Ensures computations happen on GPU if available.

### 2. **`step` Method**
```python
  def step(self, state, action, reward, next_state, done):
    self.memory.push((state, action, reward, next_state, done))  # Stores the experience in replay buffer
    self.t_step = (self.t_step + 1) % 4  # Updates step counter; learning occurs every 4 steps
    if self.t_step == 0:
      if len(self.memory.memory) > minibatch_size:
        experiences = self.memory.sample(100)  # Samples a minibatch of experiences
        self.learn(experiences, discount_factor)  # Trains the model using the sampled batch
```
- **Stores experiences in `ReplayMemory`**.
- **Limits learning to every 4 steps (`t_step % 4`)** to improve efficiency.
- **Only learns when memory has enough samples (`minibatch_size`)**.

### 3. **`act` Method**
```python
  def act(self, state, epsilon=0.):
    state = torch.from_numpy(state).float().unsqueeze(0).to(self.device)  # Converts state to tensor and moves to device
    self.local_qnetwork.eval()  # Sets network to evaluation mode (no gradient updates)
    with torch.no_grad():
      action_values = self.local_qnetwork(state)  # Gets action values from the network
    self.local_qnetwork.train()  # Resets network to training mode
    
    if random.random() > epsilon:
      return np.argmax(action_values.cpu().data.numpy())  # Selects the action with the highest value (exploitation)
    else:
      return random.choice(np.arange(self.action_size))  # Randomly selects an action (exploration)
```
- **Uses the Q-network to select an action**.
- **Implements epsilon-greedy policy**:
  - With probability `1 - epsilon`, it selects the action with the highest Q-value (exploitation).
  - With probability `epsilon`, it selects a random action (exploration).

### 4. **`learn` Method**
```python
  def learn(self, experiences, discount_factor):
    states, next_states, actions, rewards, dones = experiences  # Unpacks the experience batch
    next_q_targets = self.target_qnetwork(next_states).detach().max(1)[0].unsqueeze(1)  # Computes max Q-value for next states
    q_targets = rewards + discount_factor * next_q_targets * (1 - dones)  # Computes target Q-values using Bellman equation
    q_expected = self.local_qnetwork(states).gather(1, actions)  # Gets expected Q-values for taken actions
    loss = F.mse_loss(q_expected, q_targets)  # Calculates loss (Mean Squared Error)
    self.optimizer.zero_grad()  # Clears previous gradients
    loss.backward()  # Computes gradients via backpropagation
    self.optimizer.step()  # Updates network weights
    self.soft_update(self.local_qnetwork, self.target_qnetwork, interpolation_parameter)  # Updates target network
```
- **Updates Q-values using the Bellman equation**.
- **Computes loss using MSE**.
- **Performs backpropagation and optimizer step**.
- **Updates target network using `soft_update`**.

### 5. **`soft_update` Method**
```python
  def soft_update(self, local_model, target_model, interpolation_parameter):
    for target_param, local_param in zip(target_model.parameters(), local_model.parameters()):
      target_param.data.copy_(interpolation_parameter * local_param.data + (1.0 - interpolation_parameter) * target_param.data)
```
- **Gradually updates the target Q-network parameters**.
- **Ensures stable learning by reducing drastic changes to the target network**.

## Summary
- **The Agent class is responsible for interacting with the environment and learning from past experiences**.
- **It uses a Deep Q-Network (DQN) with experience replay and a target network for stable training**.
- **Implements an epsilon-greedy policy for balancing exploration and exploitation**.
- **Uses soft updates to gradually update the target network, preventing instability in training**.

Would you like any modifications or additional explanations? 🚀



