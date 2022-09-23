<script>
    import {
    Button,Column,Grid,Row,TextArea,Tile
    } from "carbon-components-svelte";
    import "carbon-components-svelte/css/white.css";

    let lastId = 0;

    let notes = [
        {
            id: ++lastId,
            content: "🦆 I'm not a duck.",
        },
    ];

    function createNote(ev) {
        const id = ++lastId,
            content = "";
        notes.push({ id, content });
        // Svelte compiler will only trigger an state change on variable assignment.
        notes = notes;
    }

    async function saveNote(note) {
        console.info("persisting note");
        console.info(note);
    }
</script>

<Grid>
    <Row padding="true">
        <Column>
            <Button on:click={createNote}>Create Note</Button>
        </Column>
    </Row>
    {#each notes as note}
        <Row padding="true">
            <Column>
                <Tile>
                    <TextArea value={note.content} />
                    <Button on:click={(ev) => saveNote(note)}>Save</Button>
                </Tile>
            </Column>
        </Row>
    {/each}
</Grid>
