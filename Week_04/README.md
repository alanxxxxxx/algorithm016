> BFS & DFS

* BFS 广度优先
    * 层级遍历  
    * 代码模板
    ```
        def bfs(graph, start, end)
            queue = []
            queue.append([start])
            visited.add(start)
            
            while queue:
                node = queue.pop()
                visited.add(node)
                process(node)
                nodes = generate_related_nodes(node)
                queue.push(nodes)
        
    ```
* DFS 深度优先
    * 每条分支从头至尾遍历一次
    * 代码模板
    ```
        visited = set()
        def dfs(node, visited):
            if node in visited:
                return
            visited.add(node)
            ...
            for next_node in node.children():
                if not next_node in visited:
                    dfs(next_node, visited)
    ```
    
注：每个节点遍历一次