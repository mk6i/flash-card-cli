# Flash Card CLI

<p align="center">
  <img src="https://user-images.githubusercontent.com/2894330/35425631-36b7d03a-022a-11e8-83ae-93a8cbde1d45.gif">
</p>

Flash Card CLI is a command-line study aid for MacOS. It audibly outputs two-column CSV files in sequence using the built-in `say` command, pausing between entries until the user hits the enter key to continue.

This tool is useful for language learners because it can be configured to output words using foreign voices. For instance, the tool can speak English words in the left column using an English voice and French words in the other using a French voice.

## Installation

Download the Flash Card CLI from the [releases page](https://github.com/mkaminski1988/flash-card-cli/releases) and move it to the `/usr/local/bin` directory. 

Quick install:

```bash
curl -O \
-L https://github.com/mkaminski1988/flash-card-cli/releases/download/1.0/flash-macos && \
chmod +x flash-macos && \
mv flash-macos /usr/local/bin/flash
```

## Usage

#### Input File `vocab.csv`

```csv
nageoire (f),flipper
dans le sens des aiguilles d’une montre,clockwise
reculé,"remote, distant"
calotte polaire (f),polar ice cap
vivres (nmpl),food supplies
```

> The voice engine ignores notes written in `(parentheses)`.

Read the left column using Thomas (a French voice) and the right column using Alex (an American English voice).

```bash
cat vocab.csv | flash -left Thomas -right Alex
```

The `-reverse` flag tells the speaker to read the right column before the left column.

```bash
cat vocab.csv | flash -left Thomas -right Alex -reverse
```

Mix it up! Improve retention and recall by randomly sorting multiple input files.

```
cat vocab_1.csv vocab_2.csv | sort -R | flash
```

#### Viewing available voices

MacOS offers plenty of built-in voices for multiple languages. Run the following command to see which voices are available to pass to the `-left` and `-right` flags.

```bash
say -v '?'
```

```bash
Alex                en_US    # Most people recognize me by my voice.
Alice               it_IT    # Salve, mi chiamo Alice e sono una voce italiana.
Alva                sv_SE    # Hej, jag heter Alva. Jag är en svensk röst.
Amelie              fr_CA    # Bonjour, je m’appelle Amelie. Je suis une voix canadienne.
Anna                de_DE    # Hallo, ich heiße Anna und ich bin eine deutsche Stimme.
Carmit              he_IL    # שלום. קוראים לי כרמית, ואני קול בשפה העברית.
Damayanti           id_ID    # Halo, nama saya Damayanti. Saya berbahasa Indonesia.
[..]
```
