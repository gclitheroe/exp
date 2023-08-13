import pandas as pd
from torch import Tensor
from transformers import AutoTokenizer, AutoModel
import psycopg2
from pgvector.psycopg2 import register_vector


def average_pool(last_hidden_states: Tensor,
                 attention_mask: Tensor) -> Tensor:
    last_hidden = last_hidden_states.masked_fill(~attention_mask[..., None].bool(), 0.0)
    return last_hidden.sum(dim=1) / attention_mask.sum(dim=1)[..., None]

conn = psycopg2.connect(
    host="localhost",
    user="postgres",
    password="password")

register_vector(conn)
cursor = conn.cursor()

# gte-small is trained on English only.
tokenizer = AutoTokenizer.from_pretrained("thenlper/gte-small")
model = AutoModel.from_pretrained("thenlper/gte-small")

df = pd.read_csv('data/github_issues.csv', nrows=1000000)

# generate and save embeddings for the support tickets
for ind in df.index:
    print(ind)

    # Tokenize the input texts.  gte-small max input is 512 tokens.
    batch_dict = tokenizer(df['body'][ind], max_length=512, padding=True, truncation=True, return_tensors='pt')

    outputs = model(**batch_dict)
    embeddings = average_pool(outputs.last_hidden_state, batch_dict['attention_mask'])

    cursor.execute("INSERT INTO issues (description, embedding) VALUES (%s, %s);", (df['body'][ind], embeddings.detach().numpy()[0],))
    conn.commit()


conn.close()