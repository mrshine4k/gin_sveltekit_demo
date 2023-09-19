<script>
    import Nav from "./Nav.svelte";
    import axios from "axios";
    import { onMount } from "svelte";

    //get albums on mount
    let Albums = null;
    onMount(() => {
        Albums = axios.get("http://localhost:9000/albums");
    });
</script>

<div class="card w-auto bg-base-200 overflow-hidden shadow-md p-0">
    <Nav />
</div>
<div class="card w-auto bg-base-200 shadow-md items-center m-2 p-1">
    <div class="prose">
        <h1>Welcome to SvelteKit</h1>
        <p>
            Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the
            documentation
        </p>
        <div class="overflow-x-auto max-w-xl left-1">
            <table class="table table-zebra">
                <thead>
                    <tr class="text-xl">
                        <th />
                        <th>Title</th>
                        <th>Artist</th>
                    </tr>
                </thead>
                {#await Albums}
                    <td colspan="3" align="center"><span class="loading loading-spinner loading-lg"></span></td>
                {:then Albums}
                    {#if Albums?.data}
                        <tbody>
                            {#each Albums.data as album, i}
                                <tr class="bg-base-200 hover">
                                    <td>{i + 1}</td>
                                    <td>{album.title}</td>
                                    <td>{album.artist}</td>
                                </tr>
                            {/each}
                        </tbody>
                    {/if}
                {:catch error}
                    <td colspan="3" style="color: red;" align="center">{error.message}</td>

                {/await}
            </table>
        </div>
    </div>
</div>
